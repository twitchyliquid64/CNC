var CodeGraph = (function(undefined, $) {
  var CodeGraph = function(element) {
    this.parent = element;
    this.graph = new joint.dia.Graph;
    this.paper = new joint.dia.Paper({
      el: element,
      width: 500,
      height: 200,
      model: this.graph,
      gridSize: 1
    });

    var rect = new joint.shapes.basic.Rect({
      position: {x: 100, y: 30},
      size: {width: 100, height: 30},
      attrs: { rect: { fill: 'blue' }, text: { text: 'my box', fill: 'white' } }
    });

    var rect2 = rect.clone();
    rect2.translate(300);

    var link = new joint.dia.Link({
      source: { id: rect.id },
      target: { id: rect2.id }
    });

    this.graph.addCells([rect, rect2, link]);
  }

  return CodeGraph;
})(void(0), $);
