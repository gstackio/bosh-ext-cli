package visualize

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

<table id="jamiltable" class="events-table table table-hover table-bordered" width="100%">
</table>
`

// TODO
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

const eventsJS string = `
  <script type="text/javascript">
function formatDetailsRow ( d ) {
    var context = '<div class="alert alert-success" role="alert">' +
      '<h6 class="alert-heading">Context</h6>' +
      '<pre><code>'+d["context"]+'</pre></code>' +
    '</div>';

    var error = '<div class="alert alert-danger" role="alert">' +
      '<h6 class="alert-heading">Error</h6>' +
      '<pre><code>'+d["error"]+'</pre></code>' +
    '</div>';

    var result = "";

    if (d["context"]) {
        result = result + context;
    }

    if (d["error"]) {
        result = result + error;
    }

    return result;
};

function createdRowCallback( row, data, dataIndex){
    if(data["error"]){
        $(row).addClass('list-group-item-danger');
    } else if(data["context"]) {
        $(row).addClass('list-group-item-info');
    }
};

function getUrlParameter(sParam) {
    var sPageURL = decodeURIComponent(window.location.search.substring(1)),
        sURLVariables = sPageURL.split('&'),
        sParameterName,
        i;

    for (i = 0; i < sURLVariables.length; i++) {
        sParameterName = sURLVariables[i].split('=');

        if (sParameterName[0] === sParam) {
            return sParameterName[1] === undefined ? true : sParameterName[1];
        }
    }
};

$(function() {
    $( "#filter-button" ).click(function() {
      eventsBeforeId=$( "#events-before-id" ).val();
      eventsTaskId=$( "#events-task-id" ).val();

      var url = "/events?";
      var navigate = false;
      if (eventsBeforeId) {
        url = url + "before-id=" + eventsBeforeId + "&";
        navigate = true;
      }
      if (eventsTaskId) {
        url = url + "task=" + eventsTaskId + "&";
        navigate = true;
      }

      if (navigate) {
        window.location.href = url;
      }

    });

    var taskIDFilter = getUrlParameter('task');
    var requestArguments = []
    if (taskIDFilter) {
      $( "#events-task-id" ).val(taskIDFilter);
      requestArguments.push({"name":"task","value":taskIDFilter})
    }

    var beforeIDFilter = getUrlParameter('before-id');
    if (beforeIDFilter) {
      $( "#events-before-id" ).val(beforeIDFilter);
      requestArguments.push({"name":"before-id","value":beforeIDFilter})
    }

    var depsReq = {"command":"events","arguments":requestArguments}
    $.post("/api/command", JSON.stringify(depsReq))
    .done(function(data) {
        var events = data.Tables[0].Rows

        var table = $('#jamiltable').DataTable( {
            data: events,
            columns: [
                {
                    "class":          "details-control",
                    "orderable":      false,
                    "data":           null,
                    "defaultContent": ""
                },
                { data: "id", title: "ID" },
                { data: "time", title: "Time" },
                { data: "action", title: "Action" },
                { data: "object_type", title: "Object Type" },
                { data: "object_name", title: "Object Name" },
                { data: "task_id", title: "Task ID" },
                { data: "deployment", title: "Deployment" }
                // { data: "instance", title: "Instance" }
            ],
            lengthChange: false,
            paging: false,
            ordering: false,
            buttons: [ 'pageLength'],
            createdRow: createdRowCallback
        } );

        var detailRows = [];

        $('#jamiltable tbody').on( 'click', 'tr td.details-control', function () {
            var tr = $(this).closest('tr');
            var row = table.row( tr );
            var idx = $.inArray( tr.attr('id'), detailRows );

            if ( row.child.isShown() ) {
                tr.removeClass( 'details' );
                row.child.hide();

                // Remove from the 'open' array
                detailRows.splice( idx, 1 );
            }
            else {
                tr.addClass( 'details' );
                row.child( formatDetailsRow( row.data() ) ).show();

                // Add to the 'open' array
                if ( idx === -1 ) {
                    detailRows.push( tr.attr('id') );
                }
            }
        } );

        // On each draw, loop over the detailRows array and show any child rows
        table.on( 'draw', function () {
            $.each( detailRows, function ( i, id ) {
                $('#'+id+' table.details-control').trigger( 'click' );
            } );
        } );
    });
});

  </script>
`
