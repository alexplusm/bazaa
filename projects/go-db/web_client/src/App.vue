<template>
  <div>
    "hello world!"
  </div>
</template>

<script>
import axios from "axios"

export default {
    name: 'App',
    mounted () {
        axios
            .get('/api/check/alive')
            .then(response => console.log("response: ", response))
    }
}
</script>
