<template>
  <div class="datepicker">
    <label> {{ data.data.name }}</label>
    <date-picker v-model="date" @input="getDate"></date-picker>
  </div>
</template>

<script>
import DatePicker from 'vue2-datepicker';
import 'vue2-datepicker/index.css';

export default {
  components: { DatePicker },
  props: {
    data: { type: Object }
  },
  name: "DatepickerField",
  data() {
    return {
      date: '',
    };
  },
  methods: {
    getDate() {
      // console.log('date', this.date.toISOString());
      this.$emit('getValue', {id: this.$props.data.id, label: this.date.toISOString()});
    }
  }
}
</script>

<style>
.datepicker {
  display: flex;
  flex-direction: column;
  margin-top: 15px;
}
.datepicker label {
  font-weight: 500;
  text-align: left;
}
.mx-input {
  width: 100%;
  height: 32px;
  background-color: #ffffff;
  border: 1px solid #6F8BE9;
  border-radius: 4px;
  transition: 180ms box-shadow ease-in-out;
  font-size: 16px;
  font-family: inherit;
  padding: 0.25em 0.5em;
}
.mx-input-wrapper, .mx-datepicker {
  width: 100%;
}
</style>
