import pathlib
import os
from setuptools import setup, find_packages
from pathlib import Path
from typing import List

# The directory containing this file
HERE = pathlib.Path(__file__).parent
# The text of the README file
README: str = (HERE / "README.md").read_text()


def create_file(f):
    with open(f, "w") as file:
        file.write("")
        file.close()


def get_packages(source_dir):
    create_file(os.path.join(source_dir, '__init__.py'))
    for subdir, dirs, files in os.walk(source_dir):
        create_file(os.path.join(subdir, '__init__.py'))


get_packages("gen/python")
get_packages("client/python")
packages = find_packages("gen/python")
packages2 = find_packages("client/python")
for p in packages2:
    packages.append(p)
# Setup
setup(
    name='by_email_notification',
    version="0.0.1",
    packages=["gen/python/by", "client/python/by_email_notification_client"],
    package_dir={"by_email_models": "gen/python", "by_email_client": "client/python/by_email_notification_client"},
    include_package_data=True,
    python_requires=">=3.7",
    url='https://gitlab.com/alsabor/shared/protos',
    license='LICENCESE',
    author='Fabio Moreno',
    author_email='ruberney.rodriguez@cuemby.com',
    description='gRPC models for By email Notification project',
    long_description=README,
    setup_requires=['wheel', 'twine'],
    install_requires=[
        'wheel',
        'twine',
        'grpcio>=1.23.0',
        'grpcio-tools>=1.23.0'
    ],
    classifiers=[
        "Programming Language :: Python :: 3",
        "Programming Language :: Python :: 3.6",
    ]
)
