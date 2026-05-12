// query.js 


import lib_getter_js from './getter'

////////////////////////////////////////////////////////////////////////////////


function Query_inner(q) {
    this.q = q;
}

Query_inner.prototype = {

    put_params(src) {


        let dst = this.q._params;
        if (dst == null) {
            dst = {};
            this.q._params = dst;
        }

        if (src == null) {
            return;
        }

        for (let name in src) {
            dst[name] = src[name];
        }

    },

}

////////////////////////////////////////////////////////////////////////////////


function Query_result(q, vo) {
    this.q = q;
    this.vo = vo;
}

Query_result.prototype = {

    GetItems(field_name) {
        let list = this.vo[field_name];
        if (list == null) {
            list = []
        }
        return list;
    },

    GetPagination() {

        let src = this.vo.pagination;
        let getter = lib_getter_js.NewGetter();

        let n_page = getter.From(src).Get('page').Result(1)
        let n_size = getter.From(src).Get('size').Result(1)
        let n_total = getter.From(src).Get('total').Result(0)

        let dst = {
            offset: (n_size * (n_page - 1)),
            limit: n_size,
            total: n_total,
        }
        return dst;
    },

}


////////////////////////////////////////////////////////////////////////////////

function Query() {

    this._component = null;
    this._store = null;
    this._config = {};
    this._params = {};
    this._inner = new Query_inner(this);

}

Query.prototype = {

    ForComponent: function (com) {

        let params = com.$route.query;

        this._inner.put_params(params)
        this._component = com; // not_null

        return this;
    },

    ForStore: function (store) {
        this._store = store; // not_null

        return this;
    },

    SetConfig(config) {

        if (config == null) {
            config = {};
        }

        this._inner.put_params(config.params);
        this._config = config;

        return this;
    },

    AddParams(params) {
        this._inner.put_params(params);
        return this;
    },

    GetConfig() {
        let cfg = this._config;
        if (cfg == null) {
            cfg = {}
        }
        cfg.params = this._params;
        return cfg;
    },

    GetParams() {
        return this._params;
    },

    Reload() {
        let query = this.GetParams();
        this._component.$router.push({ query })
        window.location = '#666'
    },

    HandleResultVO(vo) {
        return new Query_result(this, vo);
    },

}


////////////////////////////////////////////////////////////////////////////////

export default {

    name: "lib-query-js",

    NewQuery() {
        return new Query();
    }

}

////////////////////////////////////////////////////////////////////////////////
// EOF 
