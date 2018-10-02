package visualize

const deploymentsTemplate string = `
      <div id="deploymentsLoadingSpinner">
        <div class="row">
          <div class="col-md-6 offset-md-3" style="text-align: center;">
            <i class="fa fa-circle-o-notch fa-spin" style="font-size: 7rem"></i>
          </div>
        </div>
      </div>
      <div id="deploymentsToggle">
        <div class="row">
          <div class="col-md-6 offset-md-6" style="text-align: right;">
          	<button id="deployment-expand-all" type="button" class="btn btn-warning">Expand All</button>
			<button id="deployment-collapse-all" type="button" class="btn btn-secondary">Collapse All</button>
          </div>
        </div>
      </div>
      <div id="deploymentsSection"></div>
`
