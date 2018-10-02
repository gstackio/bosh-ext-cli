package visualize

const homeTemplate string = `

      <div class="row">
        <div class="col-xl-4 col-sm-12 mb-4">
          <div class="card text-white bg-primary o-hidden h-100">
            <div class="card-body">
              <div class="card-body-icon">
                <i class="fa fa-fw fa-tv"></i>
              </div>
              <div class="mr-5">
                <i id="deployments-count-spinner" class="fa fa-circle-o-notch fa-spin" style="font-size:24px"></i> 
                <span id="deployments-count"></span> 
                Deployments
              </div>
            </div>
            <a class="card-footer text-white clearfix small z-1" href="/deployments">
              <span class="float-left">View Details</span>
              <span class="float-right">
                <i class="fa fa-angle-right"></i>
              </span>
            </a>
          </div>
        </div>
        <div class="col-xl-4 col-sm-6 mb-4">
          <div class="card text-white bg-danger o-hidden h-100">
            <div class="card-body">
              <div class="card-body-icon">
                <i class="fa fa-fw fa-tasks"></i>
              </div>
              <div class="mr-5">
                <i id="running-tasks-count-spinner" class="fa fa-circle-o-notch fa-spin" style="font-size:24px"></i>
                <span id="running-tasks-count"></span> 
                Running Tasks
              </div>
            </div>
            <a class="card-footer text-white clearfix small z-1" href="/tasks">
              <span class="float-left">View Details</span>
              <span class="float-right">
                <i class="fa fa-angle-right"></i>
              </span>
            </a>
          </div>
        </div>
        <div class="col-xl-4 col-sm-6 mb-4">
          <div class="card text-white bg-success o-hidden h-100">
            <div class="card-body">
              <div class="card-body-icon">
                <i class="fa fa-fw fa-file"></i>
              </div>
              <div class="mr-5">
                <i id="releases-count-spinner" class="fa fa-circle-o-notch fa-spin" style="font-size:24px"></i>
                <span id="releases-count"></span> 
                Releases
              </div>
            </div>
            <a class="card-footer text-white clearfix small z-1" href="/releases">
              <span class="float-left">View Details</span>
              <span class="float-right">
                <i class="fa fa-angle-right"></i>
              </span>
            </a>
          </div>
        </div>
      </div>

`
