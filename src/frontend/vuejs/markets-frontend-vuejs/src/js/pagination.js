// pagination.js 


////////////////////////////////////////////////////////////////////////////////

function PaginationParamGetter() {


}


PaginationParamGetter.innerGetNum = function (num, min) {

    if (num == null) {
        num = min;
    }

    num = Number(num);

    if (num < min) {
        num = min;
    }

    return num;
}


PaginationParamGetter.prototype = {

    GetParams(context, dst) {

        let src = context.$route.query;

        let lim = src.limit;
        let off = src.offset;
        let fn = PaginationParamGetter.innerGetNum;

        dst.limit = fn(lim, 1);
        dst.offset = fn(off, 0);

        return dst;
    },

}


////////////////////////////////////////////////////////////////////////////////

export default {

    name: "lib-pagination-js",

    NewPaginationParamGetter() {
        return new PaginationParamGetter();
    }

}
