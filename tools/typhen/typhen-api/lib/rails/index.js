'use strict';

var path = require('path');
var fs = require('fs-extra');
var _ = require('lodash');

module.exports = function(typhen, options, helpers) {
  helpers = _.assign(helpers, {
    controllerName: function(func) {
      return typhen.helpers.upperCamelCase(func.fullName).split('::').slice(1).join('::');
    },
    controllerPath: function(symbol) {
      return typhen.helpers.underscore(symbol.fullName).split('::').slice(1).join('/');
    },
    typeName: function(type) {
      var name = type.isPrimitiveType || type.isArray ? type.name : 'TyphenApi::Model::' + type.fullName;
      return name === 'nil' ? name : typhen.helpers.upperCamelCase(name);
    }
  });

  return {
    requiredTargetModule: true,
    namespaceSeparator: '::',
    helpers: helpers,

    rename: function(symbol, name) {
      if (symbol.kind === typhen.SymbolKind.Array) {
        return 'Array[' + helpers.typeName(symbol.type) + ']';
      } else if (name === 'void') {
        return 'nil';
      }
      return name;
    },

    generate: function(g, types, modules, targetModule) {
      fs.removeSync(path.join(g.outputDirectory, 'lib/typhen_api'));

      g.generateUnlessExist('lib/rails/templates/controller/respondable.hbs', 'app/controllers/concerns/typhen_api_respondable.rb');
      g.generate('lib/rails/templates/config/initializer.hbs', 'config/initializers/typhen_api.rb');
      g.generate('lib/rails/templates/typhen_api.hbs', 'lib/typhen_api/typhen_api.rb');
      g.generate('lib/rails/templates/controller.hbs', 'lib/typhen_api/typhen_api/controller.rb');
      g.generate('lib/rails/templates/model.hbs', 'lib/typhen_api/typhen_api/model.rb');

      var functions = modules.filter(function(m) { return m === targetModule || m.ancestorModules.indexOf(targetModule) > -1; })
        .map(function(module) { return module.functions; })
        .reduce(function(a, b) { return a.concat(b); });

      g.generate('lib/rails/templates/routes.hbs', 'lib/typhen_api/typhen_api/routes.rb', functions);

      if (options.spec) {
        g.generate('lib/rails/templates/spec/support.hbs', 'spec/support/typhen_api.rb');
        g.generate('lib/rails/templates/spec/routing.hbs', 'spec/routing/generated_routing_spec.rb', functions);
      }

      functions.forEach(function(func) {
        g.generate('lib/rails/templates/controller/controller.hbs', 'underscore:lib/typhen_api/typhen_api/controller/**/*.rb', func);
        g.generateUnlessExist('lib/rails/templates/controller/app_controller.hbs', 'app/controllers/' + helpers.controllerPath(func) + '_controller.rb', func);
        g.generateUnlessExist('lib/rails/templates/controller/module.hbs', 'underscore:lib/typhen_api/typhen_api/controller/**/*.rb', func.parentModule);

        if (options.spec) {
          g.generateUnlessExist('lib/rails/templates/spec/controller_spec.hbs', 'spec/controllers/' + helpers.controllerPath(func) + '_controller_spec.rb', func);
          g.generateUnlessExist('lib/rails/templates/spec/request_spec.hbs', 'spec/requests/' + helpers.controllerPath(func) + '_spec.rb', func);
        }

        if (func.parentModule !== targetModule) {
          var modulePath = 'app/controllers/' + helpers.controllerPath(func.parentModule) + '.rb';
          g.generateUnlessExist('lib/rails/templates/controller/app_module.hbs', modulePath, func.parentModule);
        }
      });

      types.forEach(function(type) {
        switch (type.kind) {
          case typhen.SymbolKind.Enum:
            g.generate('lib/rails/templates/model/enum.hbs', 'underscore:lib/typhen_api/typhen_api/model/**/*.rb', type);
            break;
          case typhen.SymbolKind.Interface:
            if (!type.isGenericType || type.typeArguments.length > 0) {
              g.generate('lib/rails/templates/model/object.hbs', 'underscore:lib/typhen_api/typhen_api/model/**/*.rb', type);
            }
            break;
          case typhen.SymbolKind.ObjectType:
            g.generate('lib/rails/templates/model/object.hbs', 'underscore:lib/typhen_api/typhen_api/model/**/*.rb', type);
            break;
          default:
            return;
        }
        g.generateUnlessExist('lib/rails/templates/model/module.hbs', 'underscore:lib/typhen_api/typhen_api/model/**/*.rb', type.parentModule);
      });
    }
  };
};
