<template>
	<h1>loading...</h1>
</template>

<script>
export default {
layout: 'mobile',
mounted() {
    if (navigator.userAgent.match(/(phone|pad|pod|iPhone|iPod|ios|iPad|Android|Mobile|BlackBerry|IEMobile|MQQBrowser|JUC|Fennec|wOSBrowser|BrowserNG|WebOS|Symbian|Windows Phone)/i)) {
      alert("手机端");
      this.$router.replace('/mobile/index');
    } else {
      alert("pc端");
      this.$router.replace('/desktop/index');
    }
  }
}
</script>