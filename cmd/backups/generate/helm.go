package generate

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/engine"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// RenderChart Renders YAML files from 'chart/templates' directory combining Helm values from CLI switches and from .HelmValues
func (t *Templating) RenderChart(script string, gpgKeyContent string, schedule string, jobName string, image string, valuesOverride map[interface{}]interface{}) (string, error) {
	templatesDir := "cmd/backups/generate/chart/templates" // todo parametrize
	templates, templatesLoadErr := t.loadChartFiles(templatesDir)
	if templatesLoadErr != nil {
		return "", errors.Wrap(templatesLoadErr, "Cannot render Chart")
	}

	// Helm values
	values := map[string]interface{}{
		"name":          jobName,
		"scriptContent": script,
		"gpgKeyContent": gpgKeyContent,
		"schedule":      schedule,
		"image":         image,
		"scriptName":    jobName,
	}
	for key, val := range valuesOverride {
		if key == "env" {
			logrus.Infof(".HelmValues.env contains defined environment variables, if any will contain ${} or $() then will be evaluated as shell")
			var err error
			val, err = processVariablesLocally(val.(map[interface{}]interface{}))
			if err != nil {
				return "", errors.Wrap(err, "cannot parse helm environment variable list at path .HelmValues.env")
			}
		}

		values[key.(string)] = val
	}
	valuesVolume := map[string]interface{}{
		"Values": values,
	}
	// end of Helm values

	c := chart.Chart{
		Raw: nil,
		Metadata: &chart.Metadata{
			Name:        "rkc",
			Home:        "https://github.com/riotkit-org/rkc",
			Version:     "1.0",
			Description: "Backup or Restore job",
			AppVersion:  "1.0",
			Deprecated:  false,
			Type:        "application",
		},
		Lock:      nil,
		Templates: templates,
		Values:    valuesVolume,
		Schema:    nil,
		Files:     nil,
	}

	files, err := engine.Engine{Strict: true}.Render(&c, valuesVolume)
	if err != nil {
		return "", errors.Wrap(err, "cannot render a chart")
	}

	var contents []string
	for _, content := range files {
		contents = append(contents, content)
	}
	return strings.Join(contents, "\n"), nil
}

func (t *Templating) loadChartFiles(templatesDir string) ([]*chart.File, error) {
	logrus.Infof("Rendering templates from %s", templatesDir)

	files, err := ioutil.ReadDir(templatesDir)
	if err != nil {
		return []*chart.File{}, errors.Wrapf(err, "Cannot list Chart templates at path '%s'", templatesDir)
	}

	var loaded []*chart.File
	for _, f := range files {
		content, readErr := ioutil.ReadFile(templatesDir + "/" + f.Name())
		if readErr != nil {
			return []*chart.File{}, errors.Wrapf(readErr, "Cannot read Chart template at path '%s'", templatesDir+"/"+f.Name())
		}
		loaded = append(loaded, &chart.File{
			Name: f.Name(),
			Data: content,
		})
	}
	return loaded, nil
}

func processVariablesLocally(envs map[interface{}]interface{}) (map[interface{}]interface{}, error) {
	for key, value := range envs {
		if strings.Contains(value.(string), "${") || strings.Contains(value.(string), "$(") {
			byteValue, err := evaluateShell(value.(string))
			value = string(byteValue)
			envs[key] = value

			if err != nil {
				return envs, errors.Wrap(err, "cannot process helm environment variables (.HelmValues.env) in local shell")
			}
		}
	}
	return envs, nil
}

func evaluateShell(shell string) ([]byte, error) {
	logrus.Infof("Evaluating /bin/bash -c 'echo -n %s'", shell)
	c := exec.Command("/bin/bash", "-c", "echo -n "+shell)
	c.Env = os.Environ()
	return c.Output()
}
