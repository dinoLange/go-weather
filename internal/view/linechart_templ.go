// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.501
package view

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func graphLineChart(rows []RowValue, columns []Column, title string, width int, height int) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_graphLineChart_871e`,
		Function: `function __templ_graphLineChart_871e(rows, columns, title, width, height){htmx.onLoad(function(content) {
      graph(rows, columns, title, width, height)
    })

  function graph(rows, columns, title, width, height) {
    google.charts.load('current', {'packages':['corechart']});

    // Set a callback to run when the Google Visualization API is loaded.
    google.charts.setOnLoadCallback(() => drawChart(rows, columns, title, width, height) );


    // Callback that creates and populates a data table,
    // instantiates the pie chart, passes in the data and
    // draws it.
    function drawChart(rows, columns, title, width, height) {
      // Create the data table.

      var data = new google.visualization.DataTable();
      for(const column of columns) {
        data.addColumn(column.Type, column.Name)
      }
      let rowArray = []
      for(const row of rows) {
          rowArray.push([new Date(row.Time), row.Value])
      }
      
      data.addRows(rowArray);
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

      var chart = new google.visualization.LineChart(document.getElementById('line_chart'));
      chart.draw(data, options);
    }
  }}`,
		Call:       templ.SafeScript(`__templ_graphLineChart_871e`, rows, columns, title, width, height),
		CallInline: templ.SafeScriptInline(`__templ_graphLineChart_871e`, rows, columns, title, width, height),
	}
}

func LineChart(rows []RowValue, columns []Column, title string, width int, height int) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = graphLineChart(rows, columns, title, width, height).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"line_chart\"></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
