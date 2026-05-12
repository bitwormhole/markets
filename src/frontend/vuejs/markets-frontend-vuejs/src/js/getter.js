// getter.js 


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
        return p;
    },

    Reset: function () {
        this.object = {}
        this.q_path = []
        return this;
    },
}


////////////////////////////////////////////////////////////////////////////////

export default {

    name: "lib-getter-js",

    NewGetter() {
        return new Getter();
    }

}
