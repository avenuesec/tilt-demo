docker_build('tilt-demo', '.', 
    dockerfile='deployments/docker/Dockerfile')

yaml = helm('deployments/chart',
  name='tilt-demo')

k8s_yaml(yaml)