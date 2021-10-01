docker_build('tilt-demo', '.', 
    dockerfile='deployments/docker/Dockerfile')

yaml = helm('deployments/chart',
  name='tilt-demo')

k8s_yaml(yaml)

k8s_resource('btc', port_forwards=9201)
k8s_resource('quotations', port_forwards=9200)