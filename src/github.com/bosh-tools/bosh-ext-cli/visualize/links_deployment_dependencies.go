package visualize

const linksDeploymentDependenciesTemplate = `
<div id="linksConsumersDetailedLoadingSpinner">
 <div class="row">
  <div class="col-md-6 offset-md-3" style="text-align: center;">
   <i class="fa fa-circle-o-notch fa-spin" style="font-size: 7rem"></i>
  </div>
 </div>
</div>

<div id="linksConsumersDetailedContents" class="invisible">
<div class="row">
	<div class="col-md-7">
		<div class="card text-white bg-primary mb-3">
		  <div class="card-body">
			<p class="card-text">
              What deployments consume links from this deployment
            </p>
			 <div class="form-group" style="color: black;">
				<select class="form-control" id="deployment-select" data-placeholder="Choose a deployment">
				</select>
			 </div>
		  </div>
		</div>
	</div>
	<div class="col-md-5">
	</div>
</div>
 <div class="row">
	 <div class="col-md-12">
	   <div id="noChartMessage"></div>
	 </div>
 </div>
 <div id="chartHeaders" class="row invisible">
	 <div class="col-md-6">
	   <h4>Provider Deployment</h4>
	 </div>
	 <div class="col-md-6 pull-right">
	   <h4 class="pull-right">Consumer Deployments</h4>
	 </div>
 </div>
 <div class="row">
	 <div class="col-md-12">
	   <div id="chart"></div>
	 </div>
 </div>
</div>
`
