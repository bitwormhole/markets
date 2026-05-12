// licences.js 


////////////////////////////////////////////////////////////////////////////////



const the_licence_types = [
    { label: '营业执照', name: 'yyzz' },
    { label: '生产许可证(XK)', name: 'xk' },
    { label: '食品生产许可证(SC)', name: 'sc' },
];




////////////////////////////////////////////////////////////////////////////////

export default {

    name: "lib-licences-js",

    GetLicenceTypes() {
        return the_licence_types;
    }

}
