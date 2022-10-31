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
				sh 'set tname=`grep name: Chart.yaml`'
				sh 'set tversion=`grep version: Chart.yaml`'
				sh 'set hname=`echo $tname | cut -d ':' -f 2- | cut -d ' ' -f 2`'
				sh 'set hversion=`echo $tversion | cut -d ':' -f 2- | cut -d ' ' -f 2`'
				sh 'mkdir -p ~/workspace/$hname/templates'
				sh 'set chartpwd=`pwd Chart.yaml`/Chart.yaml'
				sh 'set valuepwd=`pwd values.yaml`/values.yaml'
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
