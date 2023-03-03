pipeline { 
  agent any 
 
  stages { 
      stage("Lint") {
        steps {
          catchError(buildResult: 'SUCCESS', stageResult: 'FAILURE'){
              sh "cd chesslib && golangci-lint run --enable-all ./..."
          }
        }
      }
      stage("Test") {
        steps {
          catchError(buildResult: 'SUCCESS', stageResult: 'FAILURE'){
              sh "cd chesslib/tests && go test"
          }
        }
      }
      stage("Build Docs") {
        steps {
          // catchError(buildResult: 'SUCCESS', stageResult: 'FAILURE'){
              sh "cd chesslib/resources && sh receivedoc.sh"
          // }
        }
      }
  }

  post {
    always {
        archiveArtifacts artifacts: 'chesslib/resources/docs.zip', fingerprint: true
    }
  }
}