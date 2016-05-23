graphing = (function(undefined, $) {
	exports = {}

	//=================
	//===		Graph	 ====
	//=================

	var CodeGraph = exports.CodeGraph = function(element, jsonString) {
		var self = this;

		this.parent = element;
		this.graph = new joint.dia.Graph();
		this.graph.on('add', function(cell){
			if (cell instanceof joint.dia.Link) {
				if (cell.getTargetElement() === undefined) {
					cell.remove();
				}
			}
		})

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
		});

		if (jsonString) {
			try
			{
				var json = JSON.parse(jsonString);
				this.graph.fromJSON(json);
			}
			catch(err) { console.log(err); } // Dont care about parse errors. We probs just switched to this with console.log
		}

		$(window).resize(function() {
			self.paper.setDimensions(element.innerWidth(), self.paper.getArea().height);
		})
	}

	CodeGraph.prototype.toJsonString = function() {
		return JSON.stringify(this.graph);
	}

	CodeGraph.prototype.validateConnection = function(cellViewS, magnetS, cellViewT, magnetT, end, linkView) {
		return (magnetS && magnetS.getAttribute('type') === 'output') &&
			(magnetT && magnetT.getAttribute('type') === 'input') &&
			cellViewS !== cellViewT &&
			!alreadyHasLinks(this.graph, cellViewT, magnetT);
	}

	function alreadyHasLinks(graph, cellView, magnet) {
		var port = magnet.getAttribute('port');
		var links = graph.getConnectedLinks(cellView.model, { inbound: true });

		for (var i = 0; i < links.length; i++)
			if (links[i].get('target').port == port)
				return true;

		return false;
	}

	//=====================
	//===  Joint Shapes ===
	//=====================

	// To fix defect. Source = https://groups.google.com/forum/#!topic/jointjs/md5s_fKPl_M
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
					'<text y="-2" x="3">X</text>',
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
							'<input type="text"></input>',
						'</p>',
					'</foreignObject>',
					'<g class="kill-button">',
						'<rect/>',
						'<text y="-2" x="3">X</text>',
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
					.val(this.model.prop('value'))
					.attr('placeholder', this.model.prop('placeholder'))
					.on('mousedown click', function(evt) { evt.stopPropagation(); }) // Allow the textbox to be selected
					.on('change', _.bind(function(evt) {
						this.model.prop('value', $(evt.target).val());
					}, this));
				}
			});

			//======================
			//====		blocks		====
			//======================

			CodeGraph.prototype.addCode = function(block) {
				this.graph.addCell(block.getModel());
			}

			var CodeBlock = exports.CodeBlock = function(options) {
				for (key in this.defaults) {
					this[key] = this.defaults[key];
				}

				for (key in options) {
					this[key] = options[key];
				}
			}
			CodeBlock.prototype.defaults = {'args': [], 'returns': []}

			var portHeight = 30;
			CodeBlock.prototype.getModelOptions = function() {
				return {
					size: { width: 120, height: portHeight * Math.max(this.args.length, this.returns.length)},
					position: {x: 100, y: 50},
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

			var TextCodeBlock = exports.TextCodeBlock = function(options) {
				CodeBlock.call(this, options);
			}
			TextCodeBlock.prototype = new CodeBlock({})
			TextCodeBlock.prototype.defaults = joint.util.deepSupplement({
				placeholder: 'Text'
			}, CodeBlock.prototype.defaults)

			TextCodeBlock.prototype.getModel = function() {
				var options = this.getModelOptions();
				options.size.height += 30;

				joint.util.setByPath(options, 'size/width', 150);
				joint.util.setByPath(options, 'attrs/.inPorts circle/ref-y', '0.2');
				joint.util.setByPath(options, 'attrs/.outPorts circle/ref-y', '0.2');
				joint.util.setByPath(options, 'placeholder', this.placeholder);

				return new joint.shapes.code.TextElement(options);
			}

			exports.blocks = [
				new CodeBlock({name: 'log', args: ['message']}),
				new CodeBlock({name: 'alert', args: ['message']}),
				new CodeBlock({name: 'prompt', args: ['query'], returns: ['response']}),
				new TextCodeBlock({name: 'string', returns: [''], args: ['']}),
				new TextCodeBlock({name: 'web request', returns: ['request'], placeholder: '/p/path'}),
				new CodeBlock({name: 'web write', args: ['text']}),
				new CodeBlock({name: 'email recieved', returns: ['from', 'subject', 'body']}),
			]

			return exports;
		})(void(0), $);
