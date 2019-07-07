import { Injectable } from '@angular/core';
import { HttpClient} from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})

export class ServerService {
  constructor(private http: HttpClient){}
  root = "http://localhost:6666";

  getProblem(problemid:string, userid:string){
    var turl = this.root +"/getproblem";
    var postBody = {pid: problemid, uid:userid};
		return this.http.post<Problem>(turl,postBody);
  }
  
  commit(problemid:string, userid:string){
    var turl = this.root + "/commit";
    var postBody = {pid: problemid, uid:userid};
    return this.http.post<any>(turl,postBody);
  }
}

export class Problem {
  text : string;
  time : string;
  type : string;
  tag  : string;
  try  :  number;
  ac   :  number;
  rate : string;
  leave : string; 
}
