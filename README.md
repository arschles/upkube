# upkube

A command line utility for updating Kubernetes [YAML](http://yaml.org/) manifests.

# Usage

`upkube -infile=my_yaml_file.yaml -path=some.path.in.the.yaml.file -val=valueToSetAtPath`

# Limitations

- The top level of the yaml file must be a dictionary
- Arrays are not yet supported
