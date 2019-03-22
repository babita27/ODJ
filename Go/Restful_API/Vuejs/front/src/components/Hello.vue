<template>
    <div class="friends">
       <h1 v-for="item in product" :key="item.name">
           <h4>{{ item.name }}</h4>
           <button v-on:click="decrementQty(item)">-</button>
            <input v-model="item.qty" />
            <button v-on:click="incrementQty(item)">+</button>
            <input v-model="item.price" disabled/>
        </h1>
    </div>
</template>
<script>
export default {
    data(){
        return{
            product: [
            {
                name : "Nescafe",
                price : 250,
                qty : 2,
                total: 0
            
            },
            {
                name : "Bru",
                price : 200,
                qty : 1,
                total: 0
            
            }
        ]
    }
    },
    filters:{
        ageInOneYear(age){
            return age + 1;
        },
        fullName(value){
            return `${value.first} ${value.last}`;
        }
    },
    methods:{
        incrementQty(item){
            item.qty= item.qty + 1
        },
        decrementQty(item){
            item.qty= item.qty - 1
        },
        itemTotal(item){
            item.total = item.qty * item.price;
        }
    },
    mounted(){
        fetch("http://restlearnode.academy/api/vue-5/product")
        .then(response => response.json())
        .then((data) => {
            this.product = data;
        })
    }
}
</script>