<template>
    <div id="ai">
        <div>
            <p>在浏览器中训练一个数字识别的AI模型</p>
        </div>
        <div>
            <el-row>
                <el-button :disabled=LoadDis @click="btnLoad">加载数据</el-button>
            </el-row>
            <el-row>
                <el-button :disabled=ShowDataDis @click="btnShowData">看看数据</el-button>
            </el-row>
            <el-row>
                <el-button :disabled=ShowModelDis @click="btnShowModel">看看模型</el-button>
            </el-row>
            <el-row>
                <el-button :disabled=TrainDis @click="btnTrain">训练模型</el-button>
            </el-row>
            <el-row>
                <el-button :disabled=ShowResultDis @click="btnShowResult">看结果</el-button>
            </el-row>
        </div>
        <br/><br/>
        <nuxt-link to="/desktop/goldtime">
            <el-button round type="primary">看看其它</el-button>
        </nuxt-link>
        <Comment article_id="8"/>
    </div>
</template>

<script>
    import {MnistData} from '../../tf/mnist/data';
    import * as tfvis from "@tensorflow/tfjs-vis";
    import * as tf from "@tensorflow/tfjs";
    import Comment from '../../components/desktop/Comment.vue'

    let data;
    const model = getModel();

    async function showExamples(data) {
        // Create a container in the visor
        const surface =
            tfvis.visor().surface({name: 'Input Data Examples', tab: 'Input Data'});

        // Get the examples
        const examples = data.nextTestBatch(20);
        const numExamples = examples.xs.shape[0];

        // Create a canvas element to render each example
        for (let i = 0; i < numExamples; i++) {
            //window.console.log("here");
            const imageTensor = tf.tidy(() => {
                // Reshape the image to 28x28 px
                return examples.xs
                    .slice([i, 0], [1, examples.xs.shape[1]])
                    .reshape([28, 28, 1]);
            });

            const canvas = document.createElement('canvas');
            canvas.width = 28;
            canvas.height = 28;
            canvas.style = 'margin: 4px;';
            await tf.browser.toPixels(imageTensor, canvas);
            surface.drawArea.appendChild(canvas);

            //imageTensor.dispose();
        }
    }

    function getModel() {
        const model = tf.sequential();

        const IMAGE_WIDTH = 28;
        const IMAGE_HEIGHT = 28;
        const IMAGE_CHANNELS = 1;

        model.add(tf.layers.conv2d({
            inputShape: [IMAGE_WIDTH, IMAGE_HEIGHT, IMAGE_CHANNELS],
            kernelSize: 5,
            filters: 8,
            strides: 1,
            activation: 'relu',
            kernelInitializer: 'varianceScaling'
        }));

        model.add(tf.layers.maxPooling2d({poolSize: [2, 2], strides: [2, 2]}));

        model.add(tf.layers.conv2d({
            kernelSize: 5,
            filters: 16,
            strides: 1,
            activation: 'relu',
            kernelInitializer: 'varianceScaling'
        }));
        model.add(tf.layers.maxPooling2d({poolSize: [2, 2], strides: [2, 2]}));

        model.add(tf.layers.flatten());

        const NUM_OUTPUT_CLASSES = 10;
        model.add(tf.layers.dense({
            units: NUM_OUTPUT_CLASSES,
            kernelInitializer: 'varianceScaling',
            activation: 'softmax'
        }));

        const optimizer = tf.train.adam();
        model.compile({
            optimizer: optimizer,
            loss: 'categoricalCrossentropy',
            metrics: ['accuracy'],
        });

        return model;
    }

    async function train(model, data) {
        const metrics = ['loss', 'val_loss', 'acc', 'val_acc'];
        const container = {
            name: 'Model Training', styles: {height: '1000px'}
        };
        const fitCallbacks = await tfvis.show.fitCallbacks(container, metrics);

        const BATCH_SIZE = 512;
        const TRAIN_DATA_SIZE = 5500;
        const TEST_DATA_SIZE = 1000;

        const [trainXs, trainYs] = tf.tidy(() => {
            const d = data.nextTrainBatch(TRAIN_DATA_SIZE);
            return [
                d.xs.reshape([TRAIN_DATA_SIZE, 28, 28, 1]),
                d.labels
            ];
        });

        const [testXs, testYs] = tf.tidy(() => {
            const d = data.nextTestBatch(TEST_DATA_SIZE);
            return [
                d.xs.reshape([TEST_DATA_SIZE, 28, 28, 1]),
                d.labels
            ];
        });

        return model.fit(trainXs, trainYs, {
            batchSize: BATCH_SIZE,
            validationData: [testXs, testYs],
            epochs: 10,
            shuffle: true,
            callbacks: fitCallbacks
        });
    }

    const classNames = ['Zero', 'One', 'Two', 'Three', 'Four', 'Five', 'Six', 'Seven', 'Eight', 'Nine'];

    function doPrediction(model, data, testDataSize = 500) {
        const IMAGE_WIDTH = 28;
        const IMAGE_HEIGHT = 28;
        const testData = data.nextTestBatch(testDataSize);
        const testxs = testData.xs.reshape([testDataSize, IMAGE_WIDTH, IMAGE_HEIGHT, 1]);
        const labels = testData.labels.argMax([-1]);
        const preds = model.predict(testxs).argMax([-1]);

        testxs.dispose();
        return [preds, labels];
    }

    async function showAccuracy(model, data) {
        const [preds, labels] = doPrediction(model, data);
        const classAccuracy = await tfvis.metrics.perClassAccuracy(labels, preds);
        const container = {name: 'Accuracy', tab: 'Evaluation'};
        await tfvis.show.perClassAccuracy(container, classAccuracy, classNames);

        //labels.dispose();
    }

    async function showConfusion(model, data) {
        const [preds, labels] = doPrediction(model, data);
        const confusionMatrix = await tfvis.metrics.confusionMatrix(labels, preds);
        const container = {name: 'Confusion Matrix', tab: 'Evaluation'};
        await tfvis.render.confusionMatrix(
            container, {values: confusionMatrix}, classNames);

        //labels.dispose();
    }

    async function showPrediction(model, data) {
        // Create a container in the visor
        const surface =
            tfvis.visor().surface({name: 'Prediction Examples', tab: 'Predict Data'});

        // Get the examples
        const tests = data.nextTestBatch(20);
        const numTests = tests.xs.shape[0];

        const xs = tests.xs.reshape([20, 28, 28, 1]);
        const labels = tests.labels.argMax([-1]);
        const preds = model.predict(xs).argMax([-1]);

        // Create a canvas element to render each example
        for (let i = 0; i < numTests; i++) {
            const imageTensor = tf.tidy(() => {
                // Reshape the image to 28x28 px
                return tests.xs
                    .slice([i, 0], [1, tests.xs.shape[1]])
                    .reshape([28, 28, 1]);
            });

            const canvas = document.createElement('canvas');
            canvas.width = 28;
            canvas.height = 28;
            canvas.style = 'margin: 4px;';
            await tf.browser.toPixels(imageTensor, canvas);
            surface.drawArea.appendChild(canvas);

            const canvasLabel = document.createElement('span');
            canvasLabel.innerText = preds.arraySync()[i];
            surface.drawArea.appendChild(canvasLabel)
            //imageTensor.dispose();
        }
    }

    export default {
        name: "ai",
        data() {
            return {
                LoadDis: false,
                ShowDataDis: true,
                ShowModelDis: true,
                TrainDis: true,
                ShowResultDis: true,
            }
        },
        components: {
            Comment
        },
        methods: {
            btnLoad() {
                this.LoadDis = true;
                data = new MnistData;
                data.load();
                this.ShowDataDis = false;
            },
            btnShowData() {
                this.ShowDataDis = true;
                tfvis.visor().open();
                showExamples(data);
                this.ShowModelDis = false;
            },
            btnShowModel() {
                this.ShowModelDis = true;
                tfvis.visor().open();
                tfvis.show.modelSummary({name: 'Model Architecture', tab: "Model Info"}, model);
                this.TrainDis = false;
            },
            btnTrain() {
                this.TrainDis = true;
                tfvis.visor().open();
                train(model, data);
                this.ShowResultDis = false;
            },
            btnShowResult() {
                tfvis.visor().open();
                showAccuracy(model, data);
                showConfusion(model, data);
                showPrediction(model, data)
            }
        },
    }
</script>

<style scoped>
    #ai {
        float: left;
        margin-left: 5%;
        width: 60%;
    }
</style>