import { createRouter, createWebHashHistory } from 'vue-router';
import HomePage from "@/pages/HomePage/HomePage.vue";
import CheckPage from "@/pages/CheckPage/CheckPage.vue";
import SchedulePage from "@/pages/SchedulePage/SchedulePage.vue";
import PaymentSuccessfuledPage from "@/pages/PaymentPage/PaymentSuccessfulPage.vue";
import PaymentFailedPage from "@/pages/PaymentPage/PaymentFailedPage.vue";

import test from "@/pages/testHome/testHome.vue";
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
        path: "/PaymentSuccessful",
        component: PaymentSuccessfuledPage
    }, {
        path: "/PaymentFail",
        component: PaymentFailedPage
    },{
        path: "/testHome",
        component: test
    },{
        path: "/:pathMatch(.*)",
        redirect: "/"
    }
];
const router = createRouter({
    history: createWebHashHistory(),
    routes
});
export default router;
//# sourceMappingURL=index.js.map