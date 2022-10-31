pipeline {
	environment {
		repository = 'sejunee/cicd-pipeline'

	}
    agent any
	stages {
		stage('Set Envirionment') {
			steps {
				git branch: 'master', url: 'https://github.com/qwqw1314/cicd-pipeline.git'
				sh 'mkdir -p ~/workspace/binary/'
				sh 'mkdir -p ~/workspace/$hname/templates'
				sh 'chartpwd=`pwd Chart.yaml`/Chart.yaml'
				sh 'valuepwd=`pwd values.yaml`/values.yaml'
				sh 'cp daemonset.yaml ~/workspace/$hname/templates'
			}
		}
		stage('Build') {
			steps {
				git branch: 'main', url: 'https://github.com/qwqw1314/build-image.git'
				sh 'go mod tidy'
				sh 'go build .'
				withCredentials([usernamePassword(credentialsId: 'sejunee', usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD')]) {
					sh 'docker login -u ${USERNAME} -p ${PASSWORD}'
					sh 'docker build -t $repository:latest .'
					sh 'docker push $repository:latest'
				}
			}
		}
		stage('Docker Upload') {
			steps {
				sh 'echo Docker Upload'
			}
		}
		stage('Deploy') {
			steps {
				sh 'echo Deploy'
			}
		}
	}
}
