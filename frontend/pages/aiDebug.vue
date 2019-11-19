<template>
    <div id="ai">
        <span>{{ msg }}</span>
        <div>{{ messages }}</div>
        <div>{{ images }}</div>
        <el-button @click="btnLoad">load</el-button>
        <el-button @click="btnTrain">train</el-button>
        <el-button @click="btnRun">run</el-button>
    </div>
</template>

<script>
    import {MnistData} from '@/tf/mnist/data';
    import * as model from '@/tf/mnist/model';
    import * as ui from '@/tf/mnist/ui';

    let data;
    export default {
        name: "aiDebug",
        data() {
            return {
                msg: [],
                messages: '',
                images: '',
            }
        },
        methods: {
            btnLoad() {
                data = new MnistData();
                data.load();
            },
            btnTrain() {
                model.train(data, this.msg);
            },
            btnRun() {
                const testExamples = 50;
                const batch = data.nextTestBatch(testExamples);
                const predictions = model.predict(batch.xs);
                const labels = model.classesFromLabel(batch.labels);

                this.msg = [];
                this.msg.push({batch, predictions, labels});
                this.messages = ui.showTestResults(batch, predictions, labels, this.messages, this.images);

            }
        },
        mounted() {
        }
    }
</script>

<style scoped>

</style>