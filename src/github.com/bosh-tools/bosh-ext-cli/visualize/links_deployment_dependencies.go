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
	<div class="col-md-4">
		<div class="card text-white bg-primary mb-3">
		  <div class="card-body">
			<h5 class="card-title">About this page:</h5>
			<p class="card-text">
              This page will list all the Link Consumers of the specified deployment. 
              It will also display the links established for these consumers, if any.
              Additionally, it will list the Link Providers for each consumer.   
            </p>
			 <div class="form-group" style="color: black;">
				<select class="form-control" id="deployment-select" data-placeholder="Choose a deployment">
				</select>
			 </div>
		  </div>
		</div>
	</div>
	<div class="col-md-4">
	</div>
	<div class="col-md-4">
	</div>
</div>
<div class="row">
<div class="col-md-12">
<div id="chart"></div>
</div>
</div>
</div>
`
