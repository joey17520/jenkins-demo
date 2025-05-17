pipeline {
    agent any
    
    stages {
        stage('Checkout') {
            steps {
                git branch: 'main', url: 'git@github.com:joey17520/jenkins-demo.git'
            }
        }
        
        stage('Build with Docker Compose') {
            steps {
                sh 'docker-compose down || true'
                sh 'docker-compose build --no-cache'
                sh 'docker-compose up -d'
            }
        }
    }
    
    post {
        always {
            // 清理未使用的镜像
            sh 'docker image prune -f'
        }
    }
}
