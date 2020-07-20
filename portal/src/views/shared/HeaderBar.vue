<template>
    <div id="white-header" class="header">
        <div class="header-inner-section">
            <div class="menu-icon-container">
                <img alt="Menu icon" src="@/assets/menu.png" v-on:click="toggleSidebar" />
            </div>
            <div class="text-elements">
                <div class="user-name">
                    <router-link v-if="user" :to="{ name: 'editUser' }" class="account-link">
                        {{ user.name }}
                    </router-link>
                </div>
                <div class="log-out" v-if="isAuthenticated" v-on:click="logout">Log out</div>
                <router-link :to="{ name: 'login', query: { redirect: $route.fullPath } }" class="log-in" v-if="!isAuthenticated">
                    Log in
                </router-link>
            </div>
            <UserPhoto v-if="user" :user="user" />
        </div>
    </div>
</template>

<script lang="ts">
import Vue from "vue";
import FKApi from "@/api/api";
import { mapState, mapGetters } from "vuex";
import * as ActionTypes from "@/store/actions";
import CommonComponents from "@/views/shared";
import { GlobalState } from "@/store/modules/global";

export default Vue.extend({
    name: "HeaderBar",
    components: {
        ...CommonComponents,
    },
    computed: {
        ...mapGetters({ isAuthenticated: "isAuthenticated" }),
        ...mapState({ user: (s: GlobalState) => s.user.user }),
        userImage() {
            if (this.$store.state.user.user.photo) {
                return this.$config.baseUrl + this.$store.state.user.user.photo.url;
            }
            return null;
        },
    },
    methods: {
        logout() {
            return this.$store.dispatch(ActionTypes.LOGOUT).then(() => {
                return this.$router.push({ name: "login" });
            });
        },
        toggleSidebar() {
            this.$emit("toggled");
        },
    },
});
</script>

<style scoped>
#white-header {
    width: auto;
    color: gray;
    text-align: right;
    overflow: hidden;
}
.header-inner-section {
    width: 100%;
    height: 69px;
    float: left;
    border-bottom: 2px solid #d8dce0;
}
.menu-icon-container {
    float: left;
    margin: 20px 0 0 15px;
    cursor: pointer;
}
.text-elements {
    float: right;
}
.user-name {
    padding: 12px 20px 0 0;
}
.account-link {
    text-decoration: underline;
    cursor: pointer;
}
.log-out,
.log-in {
    padding: 0 20px 0 0;
    cursor: pointer;
}
</style>