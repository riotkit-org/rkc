Backup Repository - Backup & Restore commands generator
-------------------------------------------------------

Purpose of this generator is to create procedures from templates for both **Backup** and **Restore** operations to use in automated way.
Backup made using generated procedure should be able to restore with a restore procedure in automated way.

**The generator is having two output formats:**
- shell script
- Kubernetes-like `kind: Job` and `kind: Pod`
