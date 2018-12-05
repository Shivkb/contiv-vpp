import { Injectable } from '@angular/core';
import { K8sNodeModel } from '../shared/models/k8s/k8s-node-model';
import { K8sPodModel } from '../shared/models/k8s/k8s-pod-model';
import { NodeData } from '../d3-topology/topology/topology-data/interfaces/node-data';
import { EdgeData } from '../d3-topology/topology/topology-data/interfaces/edge-data';
import { ContivDataModel } from '../shared/models/contiv-data-model';
import { VppInterfaceModel } from '../shared/models/vpp/vpp-interface-model';
import { LayoutService } from '../shared/services/layout.service';
import { TopologyType } from '../shared/interfaces/topology-type';
import { TopoColors } from '../shared/constants/topo-colors';

@Injectable({
  providedIn: 'root'
})
export class PodTopologyService {

  constructor(
    private layoutService: LayoutService
  ) { }

  public getTopologyData(data: ContivDataModel): {nodes: NodeData[], links: EdgeData[], type: TopologyType} {
    this.layoutService.podCount = {};
    const nodesTopoData = this.createNodes(data);
    const linksTopoData = this.createLinks(data);

    return {nodes: nodesTopoData, links: linksTopoData, type: 'vpp'};
  }

  private createNodes(data: ContivDataModel): NodeData[] {
    let nodesTopoData: NodeData[] = [];

    data.contivData.forEach(d => {
      const node = this.createTopologyNode(d.node);
      const pods = d.pods.map(p => this.createTopologyPod(p));
      const vswitch = this.createTopologyVswitch(d.vswitch);
      const vppPods = d.vppPods.map(p => this.createTopologyVppPod(p, vswitch));
      const bvi = this.createTopologyBVI(d.getBVI(), vswitch);

      nodesTopoData = nodesTopoData.concat([node], pods, [vswitch], vppPods, [bvi]);
    });

    return nodesTopoData;
  }

  private createLinks(data: ContivDataModel): EdgeData[] {
    const nodesLinks = this.layoutService.connectNodes(data);
    const podsLinks = this.layoutService.connectPodsToHost(data);
    const vswitchLinks = this.layoutService.connectVswitchesToHost(data);
    const vppLinks = this.layoutService.connectVppPodsToVswitch(data);
    const vxTunnels = this.layoutService.connectBVIs(data);

    return [].concat(nodesLinks, podsLinks, vswitchLinks, vppLinks, vxTunnels);
  }

  private createTopologyNode(node: K8sNodeModel): NodeData {
    const position = this.layoutService.getNodePosition(node);
    return {
      id: node.name,
      label: node.name,
      x: position.x,
      y: position.y,
      stroke: TopoColors.NODE_STROKE,
      nodeType: 'node',
      IP: node.ip
    };
  }

  private createTopologyVswitch(vswitch: K8sPodModel): NodeData {
    const position = this.layoutService.getVswitchPosition(vswitch);
    const node: NodeData = {
      id: vswitch.name,
      label: vswitch.name,
      x: position.x,
      y: position.y,
      nodeType: 'vswitch',
      IP: vswitch.podIp,
      namespace: vswitch.namespace,
      stroke: TopoColors.VSWITCH_STROKE
    };

    return node;
  }

  private createTopologyBVI(bvi: VppInterfaceModel, vswitch: NodeData): NodeData {
    const position = this.layoutService.getBVIPosition(vswitch);
    return {
      id: vswitch.label + '-bvi',
      x: position.x,
      y: position.y,
      stroke: TopoColors.BVI_STROKE,
      nodeType: 'bvi',
      IP: bvi.IPS
    };
  }

  private createTopologyPod(pod: K8sPodModel): NodeData {
    const position = this.layoutService.getPodPosition(pod);
    const node: NodeData = {
      id: pod.name,
      label: pod.name,
      x: position.x,
      y: position.y,
      nodeType: 'pod',
      IP: pod.podIp,
      namespace: pod.namespace
    };

    return node;
  }

  private createTopologyVppPod(pod: K8sPodModel, vswitch: NodeData): NodeData {
    const position = this.layoutService.getPodPosition(pod);
    const node: NodeData = {
      id: pod.name,
      label: pod.name,
      x: position.x,
      y: position.y,
      nodeType: 'vppPod',
      IP: pod.podIp,
      namespace: pod.namespace
    };

    return node;
  }

}