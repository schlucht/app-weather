import { HttpClient } from '@angular/common/http';
import { Component, inject, OnInit, signal } from '@angular/core';
import { RouterOutlet } from '@angular/router';

@Component({
  imports: [RouterOutlet],
  selector: 'app-root',
  styleUrl: './app.css',
  templateUrl: './app.html',
})
export class App implements OnInit {
  title = signal('')
  http = inject(HttpClient)
  ngOnInit(): void {
    
    this.http.post("https://fuzzy-space-disco-r76jrwwwrgr2x74r-8080.app.github.dev/user", JSON.stringify({email: "lothi@web", password: "schlucht"})).subscribe(
      (data) => {
        console.log(data)
      }
    )    
  }
}
