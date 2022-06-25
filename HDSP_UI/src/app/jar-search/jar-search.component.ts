import { Component, OnInit } from '@angular/core';
import { RestApiService } from "../shared/rest-api.service";
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-jar-search',
  templateUrl: './jar-search.component.html',
  styleUrls: ['./jar-search.component.css']
})
export class JarSearchComponent implements OnInit {
  jarData: any = [];
  jarid: any;
  temp=false;
  temp1=false;

  constructor(
    public restApi: RestApiService,
    private http: HttpClient,
    public router: Router
  ) { }

  ngOnInit() {}

  findJar() {
    {
      this.restApi.getJar(this.jarid).subscribe(data => {
        console.log("A", data);
        // console.log("B", data.length);
        this.jarData = data[0];
        // console.log("this.jarData.length", this.jarData.length)
        console.log("Type:", typeof(this.jarData))
        if(this.jarData.length != 0)
        {
          this.temp1=false;
          this.temp=true;
        }
        else
        {
          this.temp=false;
          this.temp1=true;
        }
      })
    }
  }
}
