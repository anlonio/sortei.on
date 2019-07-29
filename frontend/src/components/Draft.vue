<template>
  <v-container fluid pa-0>
    <v-layout justify-center>
      <v-flex xs12 sm8 md5>
        <v-card class="ma-4">
          <v-card-text>
            <v-layout row justify-space-around>
              <v-flex xs4>
                <v-text-field
                  box
                  v-model="min"
                  label="Mínimo"
                  :disabled="pool.length>0"
                ></v-text-field>
              </v-flex>
              <v-flex xs4>
                <v-text-field
                  box
                  v-model="max"
                  label="Máximo"
                  :disabled="pool.length>0"
                ></v-text-field>
              </v-flex>
              <v-flex xs3 md3>
                <v-btn block large class="red lighten-1" @click="clear" v-if="pool.length != 0">
                  Reiniciar
                </v-btn>
                <v-btn block large class="success" @click="initPool" v-else>
                  Iniciar
                </v-btn>
              </v-flex>
            </v-layout>
            <v-layout row v-if="poolOut.length > 0">
              <v-flex>
                <span>Anteriores:</span>
                <span v-for="n in poolOut" :key="n">
                  {{n}}
                </span>
              </v-flex>
            </v-layout>
          </v-card-text>
          <v-card-actions>
            <v-layout justify-center pb-4>
              <v-btn color="blue" large dark @click="getNumber" :disabled="pool.length == 0">Sortear</v-btn>
            </v-layout>
          </v-card-actions>
        </v-card>
      </v-flex>
    </v-layout>
    <div class="text-xs-center">
      <v-dialog v-model="dialog" width="500">
        <v-card>
          <v-card-title class="headline" primary-title>Número Sorteado</v-card-title>
          <v-card-text class="text-xs-center display-4">
            {{message}}
          </v-card-text>
          <v-divider></v-divider>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="primary" flat @click="dialog = false">Ok</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
      <v-dialog v-model="error" width="400">
        <v-card>
          <v-card-title class="headline" primary-title>Erro</v-card-title>
          <v-card-text class="title">
            {{errorMsg}}
          </v-card-text>
          <v-divider></v-divider>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="primary" flat @click="error = false">Ok</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </div>
  </v-container>
</template>

<script>
export default {
  data() {
    return {
      max:8,
      min:1,
      message: " ",
      pool:[],
      poolOut:[],
      raised: true,
      dialog: false,
      errorMsg: false,
      error: false
    };
  },
  methods: {
    getNumber() {
      let p = this.poolOut
      p.push(this.pool.pop())
      this.message = p[p.length-1]
      this.dialog = true
    },
    clear(){
      this.pool = []
      this.poolOut = []
      this.max = 8
      this.min = 1
    },
    initPool(){
      this.pool = []
      this.poolOut = []
      if(this.max <=0 || this.min <=0){
        this.errorMsg = "Digite valores válidos"
        this.error = true
      }else if(this.max <= this.min){
        this.errorMsg = "O máximo tem que ser maior que o mínimo"
        this.error = true
      }else{
        let self = this;
        window.backend.random(Number.parseInt(this.max)).then(result => {
          self.pool = result.map(r=>r+Number.parseInt(this.min))
        });
      }
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1 {
  margin-top: 2em;
  position: relative;
  min-height: 5rem;
  width: 100%;
}
a:hover {
  font-size: 1.7em;
  border-color: blue;
  background-color: blue;
  color: white;
  border: 3px solid white;
  border-radius: 10px;
  padding: 9px;
  cursor: pointer;
  transition: 500ms;
}
a {
  font-size: 1.7em;
  border-color: white;
  background-color: #121212;
  color: white;
  border: 3px solid white;
  border-radius: 10px;
  padding: 9px;
  cursor: pointer;
}
</style>
