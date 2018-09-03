package web2

const eventsTemplate string = `
<div class="row">
	<div class="col-xl-3 col-sm-6 mb-3">
		<div class="input-group sm">
		  <div class="input-group-prepend">
		    <span class="input-group-text" >Before ID</span>
		  </div>
		  <input id="events-before-id" type="text" class="input-xs form-control">
		</div>
	</div>
    <div class="col-xl-3 col-sm-6 mb-3">
    	<div class="input-group sm">
		  <div class="input-group-prepend">
		    <span class="input-group-text" >Task ID</span>
		  </div>
		  <input id="events-task-id" type="text" class="input-xs form-control">
		</div>
	</div>
	<div class="col-xl-3 col-sm-6 mb-3">
	</div>
	<div class="col-xl-3 col-sm-6 mb-3">
	</div>
</div>

<div class="row">
	<div class="col-xl-3 col-sm-6 mb-3">
	</div>
    <div class="col-xl-3 col-sm-6 mb-3">
	</div>
	<div class="col-xl-3 col-sm-6 mb-3">
	</div>
	<div class="col-xl-3 col-sm-6 mb-3">
	    <a class="btn btn-secondary btn-sm float-right" href="/events" role="button">Clear Filters</a>
		<button id="filter-button" type="button" class="btn btn-primary btn-sm float-right">Submit</button>
	</div>
</div>

<table id="jamiltable" class="events-table table table-hover table-bordered" width="100%"w>
</table>
`

// poll tasks tables every seconds
// puts releases in the hidden coulmn of the deployments
// List of releases and stemcells to upload
// get debug logs and filter them
// ace editor , use less mode color, readonly
// always poll for running tasks in a nice icon at the top
// unused releases, stemcells pie chart
// which deployment a stemcell is used in
// show deployment as locked in the list of deployment, and releases as well
// running task, use plane icon
// show debug logs of jobs
// show the duration of each event (probably can be expanded more compared to parsing the task debug logs)
