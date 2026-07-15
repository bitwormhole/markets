// getter.js 


////////////////////////////////////////////////////////////////////////////////

function convert_to_string(value) {
    if (value == null) {
        return ''
    }
    return '' + value
}

function convert_to_number(value) {
    if (value == null) {
        return 0;
    }
    return Number(value)
}

function convert_to_any(value) {
    return value // bypass
}

////////////////////////////////////////////////////////////////////////////////

function Getter() {
    this.Reset();
}

Getter.inner_sep_elist = function (elist) {

    let str = '' + elist
    let src = str.split('/')
    let dst = []

    for (let index in src) {
        let value = src[index]
        let el = value.trim()
        if (el == '') {
            continue
        }
        dst.push(el)
    }

    return dst
}

Getter.prototype = {


    /**
     * query_path: a path string like 'a/b/c/d'
    */

    Get: function (query_path) {
        this.q_path = Getter.inner_sep_elist(query_path);
        return this;
    },

    /**
     * set: target_object
    */

    From: function (target_object) {
        this.object = target_object;
        return this;
    },



    Result: function (default_value) {
        let p = this.object;
        let elist = this.q_path;
        for (let index in elist) {
            let el = elist[index];
            if (p == null) {
                break
            }
            p = p[el];
        }
        if (p == null) {
            p = default_value;
        }
        let fn = this.fn_as_t
        if (fn == null) {
            fn = convert_to_any
        }
        return fn(p);
    },

    Reset: function () {
        this.object = {}
        this.q_path = []
        this.fn_as_t = convert_to_any
        return this;
    },

    AsText() {
        return this.As(convert_to_string)
    },

    AsNumber() {
        return this.As(convert_to_number)
    },

    AsAny() {
        return this.As(convert_to_any)
    },

    As(a_type_fn) {
        if (a_type_fn != null) {
            this.fn_as_t = a_type_fn
        }
        return this
    },
}


////////////////////////////////////////////////////////////////////////////////

export default {

    name: "lib-getter-js",

    NewGetter() {
        return new Getter();
    }

}
