graphing = (function(undefined, $) {
  exports = {}

  var CodeGraph = exports.CodeGraph = function(element, jsonString) {
    var self = this;

    this.parent = element;
    this.graph = new joint.dia.Graph();
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

    if (jsonString) {
      var json = JSON.parse(jsonString);
      this.graph.fromJSON(json);
    }

    $(window).resize(function() {
      self.paper.setDimensions(element.innerWidth(), self.paper.getArea().height);
    })
  }

  CodeGraph.prototype.toJsonString = function() {
    return JSON.stringify(this.graph);
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
      size: { width: 120, height: portHeight * Math.max(this.args.length, this.returns.length)},
      code: {
        blockname: this.name
      },
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
    return new joint.shapes.code.CodeElement(this.getModelOptions());
  }

  //====================
  //===  Text blocks ===
  //====================

  joint.shapes.devs.Model.prototype.initialize = function() {
      this.updatePortsAttrs();
      this.on('change:inPorts change:outPorts', this.updatePortsAttrs, this);
      this._parent = (this._parent || this).constructor.__super__;
      this._parent.initialize.apply(this, arguments);
  }

  joint.shapes.code = {};
  joint.shapes.code.CodeElement = joint.shapes.devs.Model.extend({
      markup: [
          '<g class="rotatable">',
            '<g class="scalable">',
              '<rect class="body"/>',
            '</g>',
            '<g class="kill-button">',
              '<rect/>',
              '<text y="16" x="22.5">X</text>',
            '</g>',
            '<text class="label"/>',
            '<g class="inPorts"/>',
            '<g class="outPorts"/>',
          '</g>'].join(''),
      defaults: joint.util.deepSupplement({
          type: 'code.CodeElement'
      }, joint.shapes.devs.Model.prototype.defaults)
  });

  joint.shapes.code.CodeElementView = joint.shapes.devs.ModelView.extend({
    render: function() {
      joint.shapes.devs.ModelView.prototype.render.apply(this, arguments);

      $(this.el)
        .find('.kill-button')
        .on('mouseup click', _.bind(function(evt) {
          this.model.remove();
        }, this));
    }
  });

  // The following custom shape creates a link out of the whole element.
  joint.shapes.code.TextElement = joint.shapes.code.CodeElement.extend({
      markup: [
          '<g class="rotatable">',
            '<g class="scalable">',
              '<rect class="body"/>',
            '</g>',
            '<foreignObject>',
              '<p xmlns="http://www.w3.org/1999/xhtml">',
                '<input type="text" value="Text"></input>',
              '</p>',
            '</foreignObject>',
            '<g class="kill-button">',
              '<rect/>',
              '<text y="16" x="22.5">X</text>',
            '</g>',
            '<text class="label"/>',
            '<g class="inPorts"/>',
            '<g class="outPorts"/>',
          '</g>'].join(''),
      defaults: joint.util.deepSupplement({
          type: 'code.TextElement'
      }, joint.shapes.code.CodeElement.prototype.defaults)
  });

  joint.shapes.code.TextElementView = joint.shapes.code.CodeElementView.extend({
    render: function() {
      joint.shapes.code.CodeElementView.prototype.render.apply(this, arguments);

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

    return new joint.shapes.code.TextElement(options);
  }

  exports.blocks = [
    new CodeBlock({name: 'log', args: ['message']}),
    new CodeBlock({name: 'alert', args: ['message']}),
    new CodeBlock({name: 'prompt', args: ['query'], returns: ['response']}),
    new TextCodeBlock({name: 'string', returns: ['value']})
  ]

  return exports;
})(void(0), $);
