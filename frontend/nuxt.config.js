export default {
    mode: 'universal',
    /*
    ** Headers of the page
    */
    render: {
        resourceHints: false,
        ssr: false
    },
    head: {
        title: process.env.npm_package_name || '',
        meta: [
            {charset: 'utf-8'},
            {name: 'viewport', content: 'width=device-width, initial-scale=1'},
            {hid: 'description', name: 'description', content: process.env.npm_package_description || ''}
        ],
        link: [
            {rel: 'icon', type: 'image/x-icon', href: '/favicon.ico'}
        ]
    },
    /*
    ** Customize the progress-bar color
    */
    loading: {color: '#fff'},
    /*
    ** Global CSS
    */
    css: [
        'element-ui/lib/theme-chalk/index.css',
        '~/assets/css/main.css'
    ],
    /*
    ** Plugins to load before mounting the App
    */
    plugins: [
        '@/plugins/element-ui',
        '@/plugins/axios',
        '@/plugins/persistedstate',
    ],
    /*
    ** Nuxt.js dev-modules
    */
    buildModules: [],
    /*
    ** Nuxt.js modules
    */
    modules: [],
    /*
    ** Build configuration
    */
    build: {
        transpile: [
            /^element-ui/,
        ],
        /*
        ** You can extend webpack config here
        */
        extend(config, ctx) {
        }
    },
    env: {
        baseUrl: process.env.BASE_URL || 'http://localhost:3000'
    },
    router: {
        extendRoutes(routes) {
            routes.push({
                name: 'index',
                path: '/index',
                component: 'pages/index.vue',
            })
        }
    }
}
