import { Component, OnInit } from '@angular/core';
import {Problem,ServerService} from '../../server.service';

@Component({
  selector: 'app-commit',
  templateUrl: './commit.component.html',
  styleUrls: ['./commit.component.css']
})


export class CommitComponent implements OnInit {
 constructor(private server:ServerService) { }
 tp = new Problem;
 result = "";
 history = "";
 problemid = "";
 userid = "";
  
  ngOnInit() {
    //从url地址获取题目id
    this.problemid = "00001";
    //从cookie中获取用户id
    this.userid = "000053";
    this.getProblemText();
    this.getHistroy();
  }
  
  //将题目的详情显示到页面上
  getProblemText(){
    this.server.getProblem(this.problemid, this.userid).subscribe(result =>{
        this.tp = result;
      }
    )
  } 

  //通知后端判断上传的代码是否正确，又 题目id和用户id确定一个github 地址
  commit(){
    this.server.commit(this.problemid, this.userid).subscribe(result =>{
      this.result = result;
    }
  )
  }
  //将用户的答题状况记录显示到页面上
  getHistroy(){
    this.history = "时间： 2019-07-07 结果： 成功通过！";
  }
}
