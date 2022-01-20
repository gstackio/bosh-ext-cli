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

const homeJS string = `
  <script type="text/javascript">
    $(function() {
        depsReq = {"command":"deployments"}
      $.post("/api/command", JSON.stringify(depsReq))
          .done(function(data) {
            depsCount = data.Tables[0].Rows.length
            $("#deployments-count-spinner").remove();
            $("#deployments-count").text(depsCount);
          });

      tasksReq = {"command":"tasks","arguments":[{"name":"all"}]}
      $.post("/api/command", JSON.stringify(tasksReq))
          .done(function(data) {
            tasksCount = data.Tables[0].Rows.length
            $("#running-tasks-count-spinner").remove();
            $("#running-tasks-count").text(tasksCount);
          });

      releasesReq = {"command":"releases"}
      $.post("/api/command", JSON.stringify(releasesReq))
          .done(function(data) {
            releasesCount = data.Tables[0].Rows.length
            $("#releases-count-spinner").remove();
            $("#releases-count").text(releasesCount);
          });
    });
  </script>
`
