import shutil
import subprocess
from pathlib import Path

if __name__ == "__main__":
    if "{{ cookiecutter.git }}" == "y":
        _ = subprocess.check_call(["git", "init", "-b", "main"])
    else:
        Path(".gitignore").unlink()
        Path(".pre-commit-config.yaml").unlink()


    if "{{ cookiecutter.nix }}" == "y":
        _ = subprocess.check_call(["nix-shell", "-p", "npins", "--run", "npins init"])
    else:
        Path("shell.nix").unlink()
        shutil.rmtree(Path("npins"), ignore_errors=True)

    if "{{ cookiecutter.dockerfile }}" != "y":
        Path("Dockerfile").unlink()

    _ = subprocess.check_call(["go", "mod", "tidy"])
