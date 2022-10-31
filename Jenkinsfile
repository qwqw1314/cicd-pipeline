pipeline {
	environment {
		repository = 'sejunee/cicd-pipeline'

	}
    agent any
	stages {
		stage('Set Envirionment') {
			steps {
                sh 'mkdir -p ~/workspace/binary/'

				sh 'tname=`grep name: Chart.yaml`'
				sh 'tversion=`grep version: Chart.yaml`'
				sh 'hname=`echo $tname | cut -d ':' -f 2- | cut -d ' ' -f 2`'
				sh 'hversion=`echo $tversion | cut -d ':' -f 2- | cut -d ' ' -f 2`'
				sh 'mkdir -p ~/workspace/$hname/templates'
				sh 'chartpwd=`pwd Chart.yaml`/Chart.yaml'
				sh 'valuepwd=`pwd values.yaml`/values.yaml'
				sh 'cp daemonset.yaml ~/workspace/$hname/templates'
			}
		}
		stage('Build') {
			steps {
				git branch: 'main', url: 'https://github.com/qwqw1314/build-image.git'
				app = docker.build($repository)
				withCredentials([usernamePassword(credentialsId: 'sejunee', usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD')]) {
					sh 'docker login -u ${USERNAME} -p ${PASSWORD}'
					app.push("latest")
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
