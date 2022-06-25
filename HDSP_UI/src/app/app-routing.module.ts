import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { JarSearchComponent } from './jar-search/jar-search.component';

const routes: Routes = [
  { path: '', redirectTo: '/jardata', pathMatch: 'full' },
  { path: 'jardata', component: JarSearchComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
