package web2

const navTemplate string = `

  

  <!-- Navigation-->
  <nav class="navbar navbar-expand-lg navbar-dark bg-dark fixed-top" id="mainNav">
    <a class="navbar-brand" href="index.html">BOSH Dashboard - {{.Title}}</a>
    <button class="navbar-toggler navbar-toggler-right" type="button" data-toggle="collapse" data-target="#navbarResponsive" aria-controls="navbarResponsive" aria-expanded="false" aria-label="Toggle navigation">
      <span class="navbar-toggler-icon"></span>
    </button>
    <div class="collapse navbar-collapse" id="navbarResponsive">
      <ul class="navbar-nav navbar-sidenav" id="exampleAccordion">
        <li class="nav-item" data-toggle="tooltip" data-placement="right" title="Dashboard">
          <a class="nav-link" href="/">
            <i class="fa fa-fw fa-dashboard"></i>
            <span class="nav-link-text">Dashboard</span>
          </a>
        </li>
        <li class="nav-item" data-toggle="tooltip" data-placement="right" title="Deployments">
          <a class="nav-link" href="/deployments">
            <i class="fa fa-fw fa-area-chart"></i>
            <span class="nav-link-text">Deployments</span>
          </a>
        </li>
        <li class="nav-item" data-toggle="tooltip" data-placement="right" title="Releases">
          <a class="nav-link" href="/releases">
            <i class="fa fa-fw fa-table"></i>
            <span class="nav-link-text">Releases</span>
          </a>
        </li>
        <li class="nav-item" data-toggle="tooltip" data-placement="right" title="Tasks">
          <a class="nav-link" href="/tasks">
            <i class="fa fa-fw fa-table"></i>
            <span class="nav-link-text">Tasks</span>
          </a>
        </li>
        <li class="nav-item" data-toggle="tooltip" data-placement="right" title="Events">
          <a class="nav-link" href="/events">
            <i class="fa fa-fw fa-table"></i>
            <span class="nav-link-text">Events</span>
          </a>
        </li>
      </ul>
      <ul class="navbar-nav ml-auto">
        <li class="nav-item">
          <button id="bosh-env-btn" type="button" class="btn btn-primary" >
          </button>
        </li>
      </ul>
    </div>
  </nav>

`
