import Vue from 'vue'
import Element, {
    Aside,
    Button,
    Col,
    Container,
    Footer,
    Header,
    Main,
    Menu,
    MenuItem,
    Pagination,
    Row,
    Submenu
} from 'element-ui'
import locale from 'element-ui/lib/locale/lang/en'

Vue.use(Button);
Vue.use(Container);
Vue.use(Header);
Vue.use(Aside);
Vue.use(Main);
Vue.use(Footer);
Vue.use(Menu);
Vue.use(MenuItem);
Vue.use(Submenu);
Vue.use(Col);
Vue.use(Row);
Vue.use(locale);
Vue.use(Pagination);
Vue.use(Element, {locale});
