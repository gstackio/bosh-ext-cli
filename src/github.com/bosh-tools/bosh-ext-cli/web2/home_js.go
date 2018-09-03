package web2

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
