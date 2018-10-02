package visualize

const tasksJS string = `
<script type="text/javascript">

function createdRowCallback( row, data, dataIndex){
	if(data["state"] === "error"){
        $(row).addClass('list-group-item-danger');
    }
};

$(function() {
    var requestArguments = [
      {"name":"recent"}
    ];

	var tasksReq = {"command":"tasks","arguments":requestArguments}
    $.post("/api/command", JSON.stringify(tasksReq))
    .done(function(data) {
    	var tasks = data.Tables[0].Rows;
  
  		var table = $('#tasksTable').DataTable( {
	        data: tasks,
	        columns: [
	            { data: "id", title: "ID",
			        fnCreatedCell: function (nTd, sData, oData, iRow, iCol) {
			            $(nTd).html('<a href="/tasks-logs?task-id='+oData.id+'">'+oData.id+'</a>');
			        }
			    },
	            { data: "state", title: "State",
			        fnCreatedCell: function (nTd, sData, oData, iRow, iCol) {
			            $(nTd).html(oData.state + ' <a href="/events?task='+oData.id+'">(events)</a>');
			        }
			    },
	            { data: "started_at", title: "Started At" },
	            { data: "last_activity_at", title: "Last Activity At" },
	            { data: "user", title: "User" },
	            { data: "deployment", title: "Deployment" },
	            { data: "description", title: "Description" }
	        ],
	        lengthChange: false,
	        pageLength: 100,
	        ordering: true,
	        order: [[ 0, "desc" ]],
	        createdRow: createdRowCallback,
	        buttons: ['pageLength']
	    } );
	 
	    table.buttons().container().appendTo( '#jamiltable_wrapper .col-md-6:eq(0)' );

    });
});

</script>
`
