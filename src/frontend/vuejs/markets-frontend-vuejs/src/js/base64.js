// base64.js 


////////////////////////////////////////////////////////////////////////////////

// function ExampleObject() {
// }

// ExampleObject.prototype = {
// }

////////////////////////////////////////////////////////////////////////////////

export default {

    name: "lib-base64-js",

    Encode(plain) {
        return window.btoa(plain)
    },

    Decode(encoded) {
        return window.atob(encoded)
    },

}
