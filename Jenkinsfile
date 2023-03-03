pipeline { 
  agent any 
 
  stages { 
      stage("Lint") {
          steps {
            sh "cd chesslib"
            catchError(buildResult: 'SUCCESS', stageResult: 'FAILURE'){
                sh "golangci-lint run --enable-all ./..."
            }
          }
      }
      stage("Test") {
          steps {
            sh "cd chesslib/tests"
            catchError(buildResult: 'SUCCESS', stageResult: 'FAILURE'){
                sh "go test"
            }
          }
      }
      stage("Build Docs") {
          steps {
            sh "cd chesslib/resources"
            catchError(buildResult: 'SUCCESS', stageResult: 'FAILURE'){
                sh "sh receivedoc.sh"
            }
          }
      }
  }
}