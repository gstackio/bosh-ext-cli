package web2

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
	        pageLength: 100,
	        ordering: false,
	        buttons: [ 'pageLength'],
	        createdRow: createdRowCallback
	    } );
	 
	    table.buttons().container().appendTo( '#jamiltable_wrapper .col-md-6:eq(0)' );

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
