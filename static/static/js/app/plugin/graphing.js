graphing = (function(undefined, $) {
  exports = {}

  var CodeGraph = exports.CodeGraph = function(element) {
    var self = this;

    this.parent = element;
    this.graph = new joint.dia.Graph;
    this.paper = new joint.dia.Paper({
      el: element,
      width: element.width(),
      height: element.height(),
      model: this.graph,
      gridSize: 1,
      linkPinning: false, //Dont allow links without a target
      defaultLink: new joint.dia.Link({
        router: { name: 'manhattan' }, // Clean up the paths between links
        connector: { name: 'rounded', args: { radius: 5 }},
        attrs: { '.marker-target': { d: 'M 10 0 L 0 5 L 10 10 z' } } // Little arrow head on link endpoints
      }),
      validateConnection: function() { return self.validateConnection.apply(self, arguments); },
      snapLinks: { radius: 20 }
    });
  }

  CodeGraph.prototype.validateConnection = function(cellViewS, magnetS, cellViewT, magnetT, end, linkView) {
        lastArgs = arguments;
        // Prevent linking from input ports.
        if (magnetS && magnetS.getAttribute('type') === 'input') return false;
        // Prevent linking from output ports to input ports within one element.
        if (cellViewS === cellViewT) return false;
        // Prevent linking to input ports.
        if (!(magnetT && magnetT.getAttribute('type') === 'input')) return false;
        // Input ports cant accept more than one link
        return !alreadyHasLinks(this.graph, cellViewT, magnetT);
  }

  function alreadyHasLinks(graph, cellView, magnet) {
    var port = magnet.getAttribute('port');
    var links = graph.getConnectedLinks(cellView.model, { inbound: false });

    for (var i = 0; i < links.length; i++)
      if (links[i].get('target') == port)
        return true;

    return false;
  }

  CodeGraph.prototype.addCode = function(block) {
    this.graph.addCell(block.getModel());
  }

  var CodeBlock = exports.CodeBlock = function(options) {
    this.name = options.name;
    this.args = options.args ? options.args : [];
    this.returns = options.returns ? options.returns : [];
  }

  var portHeight = 30;
  CodeBlock.prototype.getModelOptions = function() {
    return {
      size: {width: 120, height: portHeight * Math.max(this.args.length, this.returns.length)},
      inPorts: this.args,
      outPorts: this.returns,
      attrs: {
        '.label': { text: this.name },
        '.inPorts circle': { fill: '#16A085', magnet: 'passive', type: 'input' },
        '.outPorts circle': { fill: '#E74C3C', type: 'output' },
      },
    };
  }

  CodeBlock.prototype.getModel = function() {
    return new joint.shapes.devs.Model(this.getModelOptions());
  }

  //====================
  //===  Text blocks ===
  //====================

  joint.shapes.custom = {};
  // The following custom shape creates a link out of the whole element.
  joint.shapes.custom.TextElement = joint.shapes.devs.Model.extend({
      markup: [
          '<g class="rotatable">',
            '<g class="scalable">',
              '<rect class="body"/>',
              '<foreignObject>',
                '<p xmlns="http://www.w3.org/1999/xhtml">',
                  '<input type="text" value="Text"></input>',
                '</p>',
              '</foreignObject>',
            '</g>',
            '<text class="label"/>',
            '<g class="inPorts"/>',
            '<g class="outPorts"/>',
          '</g>'].join(''),
      defaults: joint.util.deepSupplement({
          type: 'custom.TextElement'
      }, joint.shapes.devs.Model.prototype.defaults)
  });

  joint.shapes.custom.TextElementView = joint.shapes.devs.ModelView.extend({
    initialize: function() {
      joint.shapes.devs.ModelView.prototype.initialize.apply(this, arguments);
    },
    render: function() {
      joint.shapes.devs.ModelView.prototype.render.apply(this, arguments);

      $(this.el)
        .find('input')
        .on('mousedown click', function(evt) { evt.stopPropagation(); }) // Allow the textbox to be selected
        .on('change', _.bind(function(evt) {
            this.model.set('value', $(evt.target).val());
        }, this));
    }
  });

  var TextCodeBlock = exports.TextCodeBlock = function(options) {
    CodeBlock.call(this, options);
  }
  TextCodeBlock.prototype = new CodeBlock({})

  TextCodeBlock.prototype.getModel = function() {
    var options = this.getModelOptions();
    options.size.height += 30;
    options.size.width = 150;

    return new joint.shapes.custom.TextElement(options);
  }

  exports.blocks = [
    new CodeBlock({name: 'log', args: ['message']}),
    new CodeBlock({name: 'alert', args: ['message']}),
    new CodeBlock({name: 'prompt', args: ['query'], returns: ['response']}),
    new TextCodeBlock({name: 'string', returns: ['value']})
  ]

  return exports;
})(void(0), $);
