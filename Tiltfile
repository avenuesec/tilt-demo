docker_build('demo-btc', '.', 
  dockerfile='deployments/docker/Dockerfile',
  build_args={'service': 'btc'},
  only=['pkg', 'services/btc', 'go.mod', 'go.sum'],
)

docker_build('demo-quotations', '.', 
  dockerfile='deployments/docker/Dockerfile',
  build_args={'service': 'quotations'},
  only=['pkg', 'services/quotations', 'go.mod', 'go.sum'],
)

yaml = helm('deployments/chart', name='demo')

k8s_yaml(yaml)

k8s_resource('btc', port_forwards=9201)
k8s_resource('quotations', port_forwards=9200)