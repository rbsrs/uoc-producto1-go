pipeline {
    agent any

    environment {
        IMAGE_NAME = 'mi-web-uoc'
        IMAGE_TAG = 'latest'
    }

    stages {
        stage('Checkout') {
            steps {
                // Clonado del repositorio usando tu URL específica
                git branch: 'main', 
                    url: 'https://github.com/rbsrs/uoc-producto1-go.git'
            }
        }

        stage('Build Docker Image') {
            steps {
                // Construcción de la imagen Docker
                sh 'docker build -t ${IMAGE_NAME}:${IMAGE_TAG} .'
            }
        }

        stage('Deploy to Kubernetes') {
            steps {
                // Aplicación de los manifiestos de Kubernetes
                sh 'kubectl apply -f k8s/deployment.yaml'
                sh 'kubectl rollout restart deployment/mi-web-deployment'
            }
        }
    }

    post {
        success {
            echo '¡Despliegue completado con éxito!'
        }
        failure {
            echo 'El pipeline ha fallado. Revisa los logs.'
        }
    }
}
