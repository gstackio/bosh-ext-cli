package web2

const deploymentsJS string = `
  <script type="text/template" id="deploymentsList">
    <%%_.each(deployments, function(deployment, index) { %%>

      <div class="card deployment-card">
        <div class="card-header deployment-header" id="heading<%%=deployment.name%%>">
          <h5 class="mb-0">
            <button class="btn btn-info deployment-header-btn" data-toggle="collapse" data-target="#collapse<%%=deployment.name%%>">
              Dep <%%=index+1%%> - <%%=deployment.name%%>
            </button>
          </h5>
        </div>

        <div id="collapse<%%=deployment.name%%>" class="collapse deploymentSection" data-parent="#accordion">
          <div class="card-body">
            <div class="row">
              <div class="col-xl-6 col-sm-12 mb-12">
                <ul class="list-group">
                  <li class="list-group-item bg-success">Stemcells</li>

                  <%%var stemcellsList = deployment.stemcell_s.split("\n") ;%%>
                  <%%_.each(stemcellsList, function(stemcell) { %%>
                    <%%var stemcellNameVersion = stemcell.split("/") ;%%>
                    <li class="list-group-item">
                      <%%=stemcellNameVersion[0]%%> 
                      <span class="badge badge-success"><%%=stemcellNameVersion[1]%%></span> 
                    </li>
                  <%%})%%>    
                </ul>
              </div>
              <div class="col-xl-3 col-sm-6 mb-6">
                <ul class="list-group">
                  <li class="list-group-item bg-info">Configs</li>
                  <li class="list-group-item">
                    Cloud Config: 
                    <span class="badge badge-info"><%%=deployment.cloud_config%%></span>
                  </li>
                </ul>
              </div>
              <div class="col-xl-3 col-sm-6 mb-6">
                <ul class="list-group">
                  <li class="list-group-item bg-danger">Teams</li>
                  <li class="list-group-item">-</li>
                </ul>
              </div>
            </div>
            <div class="row">
              <div class="col-xl-12 col-sm-12 mb-12">
                <ul class="list-group">
                  <li class="list-group-item active">Releases Used</li>
                  <%%var releasesList = deployment.release_s.split("\n") ;%%>
                  <%%_.each(releasesList, function(release) { %%>
                    <%%var releaseNameVersion = release.split("/") ;%%>
                      <li class="list-group-item"><%%=releaseNameVersion[0]%%> <span class="badge badge-primary"><%%=releaseNameVersion[1]%%></span> </li>
                  <%%})%%>
                </ul>
              </div>
            </div>  
          </div>
        </div>
      </div>
    <%%})%%>

</script>

<script type="text/javascript">
  $(function() {        
    $( "#deployment-expand-all" ).click(function() {
      $(".deploymentSection").addClass("show");
    });

    $( "#deployment-collapse-all" ).click(function() {
      $(".deploymentSection").removeClass("show");
    });

    depsReq = {"command":"deployments"}
    $.post("/api/command", JSON.stringify(depsReq))
      .done(function(data) {
        var templateFeederTemp = {
           'deployments' : data.Tables[0].Rows
        };    
        var template = $("#deploymentsList").html();
        var toBeCompiledTemplate =_.template(template);
        var compiledTemplate = toBeCompiledTemplate(templateFeederTemp);
        $("#deploymentsLoadingSpinner").remove();
        $("#deploymentsSection").empty().html(compiledTemplate);
      });

  });
</script>
`
