import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Time } from '@angular/common';

@Injectable({
  providedIn: 'root'
})

export class ServerService {
  constructor(private http: HttpClient) { }
  root = "http://192.168.52.128:8083";

  getProblem(problemid: number, userid: string) {
    var turl = this.root + "/getproblem";
    var postBody = { pid: problemid, uid: userid };
    return this.http.post<Problem>(turl, JSON.stringify(postBody));
  }

  commit(problemid: number, userid: string) {
    var turl = this.root + "/commit";
    var postBody = { pid: problemid, uid: userid };
    return this.http.post<any>(turl, JSON.stringify(postBody));
  }
}

export class Problem {
  id: number
  title: string
  description: string
  type: string
  code_type: string
  checkout_path: string
  attach_code: string
  attach_file: string
  answer: string
  status: boolean
  createTime: Time
}
