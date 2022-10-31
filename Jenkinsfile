pipeline {
	environment {
		repository = 'sejunee/cicd-pipeline'
		kubeconfig = credentials('kubeconfig')
	}
    agent any
	stages {
		stage('Set Environment') {
			steps {
				echo 'set env'
				git branch: 'master', url: 'https://github.com/qwqw1314/cicd-pipeline.git'
				sh 'mkdir -p ~/workspace/binary/'
				sh 'cp Chart.yaml values.yaml ~/workspace/' 
				sh 'cp daemonset.yaml ~/workspace/$hname/templates'
			}
		}
		stage('Build') {
			steps {
				echo 'build'
				git branch: 'main', url: 'https://github.com/qwqw1314/build-image.git'
				sh 'go mod tidy'
				sh 'go build .'
				withCredentials([usernamePassword(credentialsId: 'sejunee', usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD')]) {
					sh 'docker login -u ${USERNAME} -p ${PASSWORD}'
					sh 'docker build -t $repository:latest .'
				}
			}
		}
		stage('Docker Upload') {
			steps {
				echo 'image upload'
				sh 'docker push $repository:latest'
			}
		}
		stage('Helm Initializing') {
			steps {
				echo 'helm init'
				sh 'cd ~/workspace'
                sh 'helm create ~/workspace/daemonset'
                sh 'cp ~/workspace/Chart.yaml ~/workspace/values.yaml ~/workspace/daemonset/'
                sh 'cd ~/workspace/daemonset/templates'
				dir("../daemonset/templates") {
	                sh 'rm -rf `ls | grep -v daemonset.yaml`'
					sh 'helm lint ../'
				}
				script {
					HELM_EXIST = sh (
						script: 'helm list --kubeconfig=${kubeconfig} | grep daemonset',
						returnStdout: true
					)
				}
			}
		}
		stage('Helm Install') {
			when {
				expression {
					return ${HELM_EXIST} == '';
				}
			}
			steps {
				echo 'helm install'

			}
		}
		stage('Helm Upgrade') {
			when {
				expression {
					return ${HELM_EXIST} != '';
				}
			}
			steps {
				echo 'helm upgrade'

			}
		}
		stage('Cleanup') {
			steps {
				echo 'cleanup'
			}
		}
	}
}
