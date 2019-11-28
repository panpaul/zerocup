<template>
    <div id="reader">
        <audio :src="sound" id="bgAudio" style="display: none"/>
        <mt-palette-button :radius="100" class="pb" content="+" direction="lt" ref="target">
            <div @touchstart="GoToTop(1)" class="blank-button icon-popup">
                <img alt="up" src="../../static/arrow.png" style="width: 50px"/>
            </div>
            <div @touchstart="Read()" class="blank-button icon-popup">
                <img alt="up" src="../../static/sound.png" style="width: 50px" v-if="!playStatus"/>
                <img alt="up" src="../../static/stop.png" style="width: 50px" v-if="playStatus"/>
            </div>

        </mt-palette-button>
    </div>
</template>

<script>
    export default {
        name: "read",
        data() {
            return {
                playStatus: false,
            }
        },
        props: {
            sound: String
        },
        watch: {
            sound: function () {
                if (this.playStatus) {
                    this.playStatus = false;
                }
            }
        },
        methods: {
            GoToTop(val) {
                //console.log(val);
                this.$refs.target.collapse();
                window.scrollTo(0, 0);
            },
            Read() {
                if (this.playStatus) {
                    document.getElementById('bgAudio').pause();
                } else {
                    document.getElementById('bgAudio').play();
                }
                this.playStatus = !this.playStatus;
                this.$refs.target.collapse();
            }
        }
    }
</script>

<style scoped>
    .pb {
        width: 70px;
        height: 70px;
        line-height: 70px;
        color: #FFF;
        position: fixed;
        bottom: 75px;
        right: 30px;
    }

    .icon-button {
        width: 50px;
        height: 50px;
        border-radius: 50%;
        background-color: #26a2ff;
        color: #fff;
        line-height: 50px;
        text-align: center;
    }

    .blank-button {
        width: 50px;
        height: 50px;
    }

</style>