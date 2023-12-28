class LineChart extends HTMLElement {
  // The browser calls this method when the element is
  // added to the DOM.

  columns
  rows
  title
  width
  height

  static get observedAttributes() {
    return ['columns', 'rows', 'title', 'width', 'height'];
  }

  attributeChangedCallback(name, oldValue, newValue) {
    // Called when one of the observed attributes changes
    if (name === 'columns') {
      this.columns = this.getAttribute('columns');
    }
    if (name === 'rows') {
      this.rows = JSON.parse(this.getAttribute('rows'));
    }
    if (name === 'title') {
      this.title = this.getAttribute('title');
    }
    if (name === 'width') {
      this.width = this.getAttribute('width');
    }
    if (name === 'height') {
      this.height = this.getAttribute('height');
    }
  }


  connectedCallback() {

    google.charts.load('current', {'packages':['corechart']});

    // Set a callback to run when the Google Visualization API is loaded.
    google.charts.setOnLoadCallback(() => drawChart(this, this.columns, this.rows, this.title, this.width, this.height));


    // Callback that creates and populates a data table,
    // instantiates the pie chart, passes in the data and
    // draws it.
    function drawChart(element, columns, rows, title, width, height) {
      // Create the data table.

      var data = new google.visualization.DataTable();
      var columnArray = columns.split(",")
      for(let index in columnArray) {
        data.addColumn(columnArray[index].split(":")[0].trim(), columnArray[index].split(":")[1].trim())
      }
      data.addRows(rows);
      if (!title) {
        title = 'Default title'
      }
      if (!width) {
        width = '400'
      }
      if (!height) {
        height = '300'
      }
      const options = {
          title: title,
          width: width,
          height: height
      }

        var chart = new google.visualization.LineChart(element);
        chart.draw(data, options);
    }

  }
}

// Register the CurrentDate component using the tag name <current-date>.
customElements.define('line-chart', LineChart);