import { createRouter, createWebHashHistory } from 'vue-router';
import HomePage from "@/pages/HomePage/HomePage.vue";
import CheckPage from "@/pages/CheckPage/CheckPage.vue";
import SchedulePage from "@/pages/SchedulePage/SchedulePage.vue";
import PaymentPage from "@/pages/PaymentPage/PaymentPage.vue";
import PaymentFailedPage from "@/pages/PaymentPage/PaymentFailedPage.vue";

import testQRCode from "@/pages/testHome/testQRCode.vue";
const routes = [{
        path: "/",
        component: HomePage
    }, {
        path: "/Schedule",
        component: SchedulePage
    }, {
        path: "/Check",
        component: CheckPage
    }, {
        name:'PaymentPage',
        path: '/PaymentPage',
        component: PaymentPage,
}, {
        path: "/PaymentFail",
        component: PaymentFailedPage
    }, {
    path: "/test",
    component: testQRCode
    }
    // },{
    //     path: "/:pathMatch(.*)",
    //     redirect: "/"
    // }
];
const router = createRouter({
    history: createWebHashHistory(),
    routes
});
export default router;
//# sourceMappingURL=index.js.map